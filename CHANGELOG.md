CHANGELOG
---------
- **2018-01-28**
  - Deprecate `1.7.6` and `1.6.4`: `../../pkg/sftp/packet-manager.go:44: undefined: sort.Slice`
- **2017-11-17**
  - Add `jsonutil.PrettyPrint()`
  - Update `urlutil.ToSlug()`
- **2017-11-02**
  - Add `timeutil.XOXTimes` `timeutil.YOYTimes()` `timeutil.QOQTimes()`
  - Add `sortutil.Uint16s()`
- **2017-10-30**
  - Add `mathutil.RangeInt64` and `mathutil.RangeFloat64`
  - Add `timeutil.MinMax()` and `timeutil.RFC3339YMDTime`
- **2017-09-27**
  - Add `raymondhelpers.Register()`
  - Add `reflectutil.GetField()`
  - Add `template.Execute()`
  - Add `timeutil.GetFormat()`
- **2017-09-24**
  - Add `phonenumber.AreaCodeToGeo`
  - Add `timeutil.TimeRange`
- **2017-09-20**
  - Add `multipartutil.MultipartBuilder`
  - Add `urlutil.URLInfo`
  - Add `errorsutil.PanicOnErr`
  - Change `httputil` to `httputilmore`
- **2017-09-19**
  - Add `httputil.RateLimitInfo`
  - Add `httputil.NewResponseRateLimitInfo`
  - Add `strconvtuil.AtoiWithDefault`
- **2017-07-08**
  - Add `stringsutil.JoinInterface`
- **2016-11-19**
  - Add `stringsutil.CondenseString`
- **2016-11-13**
  - Add `timeutil.ISO8601Z` time format
- **2016-08-13**
  - Add `timeutil.FromTo` time format