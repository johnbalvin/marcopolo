package trace

import (
	"errors"
)

// ErrEmpty error to tell the fiels are empty
var ErrEmpty = errors.New("err Empty restult")
var ErrNotEmpty = errors.New("err Not Empty restult")

// ErrMaxAttempt error to tell the max attemps has reached
var ErrMaxAttempt = errors.New("err Max attemps")
var ErrHashStatus = errors.New("err public file but not hash or raw pdf presented")

// ErrMaxAttempt error to tell the max attemps has reached
var ErrStatusCode = errors.New("err Not a correct status code")
var ErrDuplicated = errors.New("err  Duplicated")
var ErrDeleted = errors.New("err Deleted")
var ErrParameter = errors.New("err not correct parameters")
var ErrCounty = errors.New("err not correct county")
var ErrParallellogin = errors.New("err not paralle login")
var ErrPlan = errors.New("err not correct plan")
var ErrCookies = errors.New("err not correct cookies")
var ErrConfution = errors.New("err ErrConfution")
var ErrLogin = errors.New("err while login")
var ErrNotEqueal = errors.New("err parameters not equal")
var ErrLimit = errors.New("err, limit reached")
var ErrPermissions = errors.New("err not correct permissions")
var ErrCaptcha = errors.New("err not correct captcha")
var ErrWaf = errors.New("err not correct captcha")
var ErrSuggestions = errors.New("err, multiple results")

var ErrCaptchaUnsolved = errors.New("err  captcha unsolvable")
var ErrWrongCaptcha = errors.New("err wrong captcha")
var ErrNotSchema = errors.New("not correct schema")
var ErrMaxQuota = errors.New("max quota")
var ErrSiteDown = errors.New("site down")
var ErrFormat = errors.New("not correct format")
var ErrCaptchaZeroBalance = errors.New("Captcha, Zero balance")
