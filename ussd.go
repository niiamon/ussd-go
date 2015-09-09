package ussd

import (
	"log"
	"reflect"
	"regexp"

	"github.com/samora/ussd-go/sessionstores"
)

// Action func
type Action func() Response

// Middleware func
type Middleware func(*Context)

// Data map
type Data map[string]interface{}

type route struct {
	Ctrl, Action string
}

// Request from USSD.
type Request struct {
	Mobile  string
	Message string
	Network string
}

// Response to USSD.
type Response struct {
	Message           string
	Release, redirect bool
	err               error
	route             route
}

// Ussd sets up USSD.
type Ussd struct {
	initialRoute    route
	session         *session
	store           sessionstores.Store
	middlewares     []Middleware
	ctrls           map[string]interface{}
	context         *Context
	initiationRegex *regexp.Regexp
}

// New USSD
func New(store sessionstores.Store, ctrl, action string) *Ussd {
	u := &Ussd{
		initialRoute:    route{StrTrim(ctrl), StrTrim(action)},
		store:           store,
		middlewares:     make([]Middleware, 0),
		ctrls:           make(map[string]interface{}),
		initiationRegex: regexp.MustCompile(`^\*\d+[\*|#]`),
	}
	u.Ctrl(new(core))
	return u
}

// Middleware registers a middleware function.
// Middlwares are executed in order of addition.
// Middlwares are executed before an action.
// Middlewares are executed once per request.
func (u *Ussd) Middleware(m Middleware) {
	u.middlewares = append(u.middlewares, m)
}

// Ctrl registers a controller that has related actions.
func (u *Ussd) Ctrl(c interface{}) {
	name := reflect.ValueOf(c).Elem().Type().Name()
	if name == StrEmpty {
		panicln("ussd: Ctrl only accepts named types")
	}
	if _, ok := u.ctrls[name]; ok {
		panicln("ussd: %v ctrl already exists", name)
	}
	u.ctrls[name] = c
}

// Process USSD request.
func (u *Ussd) process(request *Request) *Response {
	uCopy := new(Ussd)
	*uCopy = *u
	err := uCopy.store.Connect()
	if err != nil {
		log.Panicln(err)
	}
	defer uCopy.store.Close()
	request.Network = StrLower(request.Network)
	request.Message = StrTrim(request.Message)

	// setup context
	uCopy.context = new(Context)
	uCopy.context.DataBag = newDataBag(uCopy.store, request)
	uCopy.context.Data = make(Data)
	uCopy.context.Request = request

	// setup session
	uCopy.session = newSession(uCopy.store, uCopy.context.Request)

	// execute middlewares
	for _, m := range uCopy.middlewares {
		m(uCopy.context)
	}

	return uCopy.exec()
}

// ProcessWithAdapters processes USSD using adapters
func (u *Ussd) ProcessWithAdapters(request RequestAdapter, response ResponseAdapter) {
	res := u.process(request.GetRequest())
	response.SetResponse(res)
}

// ProcessSmsgh processes USSD from SMSGH
func (u *Ussd) ProcessSmsgh(request *SmsghRequest) *SmsghResponse {
	response := new(SmsghResponse)
	u.ProcessWithAdapters(request, response)
	return response
}

// ProcessNsano processes USSD from Nsano
func (u *Ussd) ProcessNsano(request *NsanoRequest) *NsanoResponse {
	response := new(NsanoResponse)
	u.ProcessWithAdapters(request, response)
	return response
}

func (u *Ussd) exec() *Response {
	response := new(Response)
	if u.context.Request.Message == "" {
		u.end()
		return response
	}
	if u.initiationRegex.MatchString(u.context.Request.Message) == true {
		return u.onInitiation()
	}
	return u.onResponse()
}

func (u *Ussd) onInitiation() *Response {
	u.end()
	r := route{u.initialRoute.Ctrl, u.initialRoute.Action}
	u.session.Set(r)
	return u.onResponse()
}

func (u *Ussd) onResponse() *Response {
	for {
		exists := u.session.Exists()
		if !exists {
			panicln("ussd: User %v's session not found",
				u.context.Request.Mobile)
		}
		r := u.session.Get()

		res := u.execHandler(r)
		if res.err != nil {
			log.Println(res.err)
			u.end()
			return res
		}
		if res.redirect {
			r = route{res.route.Ctrl, res.route.Action}
			u.session.Set(r)
			continue
		}
		if !res.Release {
			u.session.Set(res.route)
		}
		return res
	}
}

func (u *Ussd) end() {
	u.context.DataBag.Clear()
	u.session.Close()
}

func (u *Ussd) execHandler(r route) *Response {
	c, ok := u.ctrls[r.Ctrl]
	if !ok {
		panicln("ussd: %v ctrl not found", r.Ctrl)
	}

	m, ok := reflect.TypeOf(c).MethodByName(r.Action)
	if !ok {
		panicln("ussd: %v has no action %v", r.Ctrl, r.Action)
	}
	args := []reflect.Value{
		reflect.ValueOf(c), reflect.ValueOf(u.context)}
	rv := m.Func.Call(args)[0]
	res, ok := rv.Interface().(*Response)
	if !ok {
		panicln("ussd: %v action on %v ctrl must return Response",
			r.Ctrl, r.Action)
	}
	return res
}
