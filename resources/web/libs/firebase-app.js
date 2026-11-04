// No-op stand-in for the Firebase SDK: analytics are stripped from this build.
// The compiled Flutter app skips loading Firebase from www.gstatic.com when `firebase`
// already exists, and only uses initializeApp / app / analytics; every analytics call
// becomes a no-op, so nothing is sent anywhere.
(function () {
  var noop = function () {
    return Promise.resolve();
  };
  var analytics = new Proxy(
    {},
    {
      get: function () {
        return noop;
      },
    }
  );
  var apps = [];
  function initializeApp(options, name) {
    name = name || "[DEFAULT]";
    for (var i = 0; i < apps.length; i++) {
      if (apps[i].name === name) return apps[i];
    }
    var app = {
      name: name,
      options: options || {},
      automaticDataCollectionEnabled: false,
      delete: noop,
    };
    apps.push(app);
    return app;
  }
  window.firebase = {
    SDK_VERSION: "8.10.1",
    apps: apps,
    initializeApp: initializeApp,
    app: function (name) {
      return initializeApp({}, name);
    },
    analytics: function () {
      return analytics;
    },
  };
  window.firebase.analytics.isSupported = function () {
    return Promise.resolve(false);
  };
})();
