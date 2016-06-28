'use strict';

var express = require('express');
var kraken = require('kraken-js');


var options, app;

/*
 * Create and configure application. Also exports application instance for use by tests.
 * See https://github.com/krakenjs/kraken-js#options for additional configuration options.
 */
options = {
    onconfig: function (config, next) {
        /*
         * Add any additional config setup or overrides here. `config` is an initialized
         * `confit` (https://github.com/krakenjs/confit/) configuration object.
         */
        next(null, config);
    }
};

app = module.exports = express();
app.use(kraken(options));

// session相关,将session存储到mongodb
const session = require('express-session');
const MongoStore = require('connect-mongo')(session);
app.use(session({
    secret: 'iqjmvh-178fd-fwh8f-cfenp',
    resave: true,
    saveUninitialized: true,
    store: new MongoStore({url: 'mongodb://192.168.1.119:27017/test'})
}));

app.on('start', function () {
    console.log('Application ready to serve requests.');
    console.log('Environment: %s', app.kraken.get('env:env'));
});
