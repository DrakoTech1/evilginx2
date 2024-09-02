// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package zap provides fast, structured, leveled logging.
//
// For applications that log in the hot path, reflection-based serialization
// and string formatting are prohibitively expensive - they're CPU-intensive
// and make many small allocations. Put differently, using json.Marshal and
// fmt.Fprintf to log tons of interface{} makes your application slow.
//
// Zap takes a different approach. It includes a reflection-free,
// zero-allocation JSON encoder, and the base Logger strives to avoid
// serialization overhead and allocations wherever possible. By building the
// high-level SugaredLogger on that foundation, zap lets users choose when
// they need to count every allocation and when they'd prefer a more familiar,
// loosely typed API.
//
<<<<<<< HEAD
// # Choosing a Logger
=======
// Choosing a Logger
>>>>>>> deathstrox/main
//
// In contexts where performance is nice, but not critical, use the
// SugaredLogger. It's 4-10x faster than other structured logging packages and
// supports both structured and printf-style logging. Like log15 and go-kit,
// the SugaredLogger's structured logging APIs are loosely typed and accept a
// variadic number of key-value pairs. (For more advanced use cases, they also
// accept strongly typed fields - see the SugaredLogger.With documentation for
// details.)
<<<<<<< HEAD
//
//	sugar := zap.NewExample().Sugar()
//	defer sugar.Sync()
//	sugar.Infow("failed to fetch URL",
//	  "url", "http://example.com",
//	  "attempt", 3,
//	  "backoff", time.Second,
//	)
//	sugar.Infof("failed to fetch URL: %s", "http://example.com")
=======
//  sugar := zap.NewExample().Sugar()
//  defer sugar.Sync()
//  sugar.Infow("failed to fetch URL",
//    "url", "http://example.com",
//    "attempt", 3,
//    "backoff", time.Second,
//  )
//  sugar.Infof("failed to fetch URL: %s", "http://example.com")
>>>>>>> deathstrox/main
//
// By default, loggers are unbuffered. However, since zap's low-level APIs
// allow buffering, calling Sync before letting your process exit is a good
// habit.
//
// In the rare contexts where every microsecond and every allocation matter,
// use the Logger. It's even faster than the SugaredLogger and allocates far
// less, but it only supports strongly-typed, structured logging.
<<<<<<< HEAD
//
//	logger := zap.NewExample()
//	defer logger.Sync()
//	logger.Info("failed to fetch URL",
//	  zap.String("url", "http://example.com"),
//	  zap.Int("attempt", 3),
//	  zap.Duration("backoff", time.Second),
//	)
=======
//  logger := zap.NewExample()
//  defer logger.Sync()
//  logger.Info("failed to fetch URL",
//    zap.String("url", "http://example.com"),
//    zap.Int("attempt", 3),
//    zap.Duration("backoff", time.Second),
//  )
>>>>>>> deathstrox/main
//
// Choosing between the Logger and SugaredLogger doesn't need to be an
// application-wide decision: converting between the two is simple and
// inexpensive.
<<<<<<< HEAD
//
//	logger := zap.NewExample()
//	defer logger.Sync()
//	sugar := logger.Sugar()
//	plain := sugar.Desugar()
//
// # Configuring Zap
=======
//   logger := zap.NewExample()
//   defer logger.Sync()
//   sugar := logger.Sugar()
//   plain := sugar.Desugar()
//
// Configuring Zap
>>>>>>> deathstrox/main
//
// The simplest way to build a Logger is to use zap's opinionated presets:
// NewExample, NewProduction, and NewDevelopment. These presets build a logger
// with a single function call:
<<<<<<< HEAD
//
//	logger, err := zap.NewProduction()
//	if err != nil {
//	  log.Fatalf("can't initialize zap logger: %v", err)
//	}
//	defer logger.Sync()
=======
//  logger, err := zap.NewProduction()
//  if err != nil {
//    log.Fatalf("can't initialize zap logger: %v", err)
//  }
//  defer logger.Sync()
>>>>>>> deathstrox/main
//
// Presets are fine for small projects, but larger projects and organizations
// naturally require a bit more customization. For most users, zap's Config
// struct strikes the right balance between flexibility and convenience. See
// the package-level BasicConfiguration example for sample code.
//
// More unusual configurations (splitting output between files, sending logs
// to a message queue, etc.) are possible, but require direct use of
// go.uber.org/zap/zapcore. See the package-level AdvancedConfiguration
// example for sample code.
//
<<<<<<< HEAD
// # Extending Zap
=======
// Extending Zap
>>>>>>> deathstrox/main
//
// The zap package itself is a relatively thin wrapper around the interfaces
// in go.uber.org/zap/zapcore. Extending zap to support a new encoding (e.g.,
// BSON), a new log sink (e.g., Kafka), or something more exotic (perhaps an
// exception aggregation service, like Sentry or Rollbar) typically requires
// implementing the zapcore.Encoder, zapcore.WriteSyncer, or zapcore.Core
// interfaces. See the zapcore documentation for details.
//
// Similarly, package authors can use the high-performance Encoder and Core
// implementations in the zapcore package to build their own loggers.
//
<<<<<<< HEAD
// # Frequently Asked Questions
=======
// Frequently Asked Questions
>>>>>>> deathstrox/main
//
// An FAQ covering everything from installation errors to design decisions is
// available at https://github.com/uber-go/zap/blob/master/FAQ.md.
package zap // import "go.uber.org/zap"
