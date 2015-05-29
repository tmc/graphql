/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _GraphQLWebClientWrapper = __webpack_require__(2);

	var _GraphQLWebClientWrapper2 = _interopRequireDefault(_GraphQLWebClientWrapper);

	_react2['default'].render(_react2['default'].createElement(_GraphQLWebClientWrapper2['default'], { endpoint: 'http://tmc.parseapp.com/graphql' }), document.getElementById('root'));

/***/ },
/* 1 */
/***/ function(module, exports, __webpack_require__) {

	module.exports = React;

/***/ },
/* 2 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	var _get = function get(_x, _x2, _x3) { var _again = true; _function: while (_again) { var object = _x, property = _x2, receiver = _x3; desc = parent = getter = undefined; _again = false; var desc = Object.getOwnPropertyDescriptor(object, property); if (desc === undefined) { var parent = Object.getPrototypeOf(object); if (parent === null) { return undefined; } else { _x = parent; _x2 = property; _x3 = receiver; _again = true; continue _function; } } else if ('value' in desc) { return desc.value; } else { var getter = desc.get; if (getter === undefined) { return undefined; } return getter.call(receiver); } } };

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _GraphQLWebClient = __webpack_require__(3);

	var _GraphQLWebClient2 = _interopRequireDefault(_GraphQLWebClient);

	var styles = {
	  clear: { clear: 'both' }
	};

	var GraphQLWebClientWrapper = (function (_React$Component) {
	  function GraphQLWebClientWrapper(props) {
	    _classCallCheck(this, GraphQLWebClientWrapper);

	    _get(Object.getPrototypeOf(GraphQLWebClientWrapper.prototype), 'constructor', this).call(this, props);
	    var endpoint = this.props.endpoint;
	    if (window && window.location.search) {
	      var m = window.location.search.match(/endpoint=(.+)/)[1];
	      if (m) {
	        endpoint = m;
	      }
	    }
	    this.state = {
	      endpoint: endpoint,
	      cannedQueries: ['{ __schema { root_fields { name, description } } }', '{ __types { name, description} }', '{ __types { name, description, fields { name, description } } }', '{ _User { __type__ { fields { name } } } }']
	    };
	    this.state.defaultQuery = this.state.cannedQueries[0];
	    if (window.location.hash.length > 1) {
	      this.state.defaultQuery = decodeURIComponent(window.location.hash.slice(1));
	    }
	  }

	  _inherits(GraphQLWebClientWrapper, _React$Component);

	  _createClass(GraphQLWebClientWrapper, [{
	    key: 'onChange',
	    value: function onChange(event) {
	      this.setState({ endpoint: event.target.value });
	    }
	  }, {
	    key: 'onCannedQueryClicked',
	    value: function onCannedQueryClicked(event) {
	      this.setState({ defaultQuery: event.target.text });
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      var _this = this;

	      var cannedQueries = this.state.cannedQueries.map(function (query) {
	        return _react2['default'].createElement(
	          'li',
	          null,
	          _react2['default'].createElement(
	            'a',
	            { href: '#', onClick: _this.onCannedQueryClicked.bind(_this) },
	            query
	          )
	        );
	      });
	      return _react2['default'].createElement(
	        'div',
	        null,
	        _react2['default'].createElement(
	          'h1',
	          null,
	          'graphql client'
	        ),
	        _react2['default'].createElement(
	          'label',
	          null,
	          'graphql endpoint:'
	        ),
	        _react2['default'].createElement('input', { size: '50', defaultValue: this.state.endpoint, onChange: this.onChange.bind(this) }),
	        _react2['default'].createElement('hr', null),
	        _react2['default'].createElement(_GraphQLWebClient2['default'], {
	          defaultQuery: this.state.defaultQuery,
	          endpoint: this.state.endpoint
	        }),
	        _react2['default'].createElement(
	          'ul',
	          { style: styles.clear },
	          _react2['default'].createElement(
	            'li',
	            null,
	            'Canned Queries:'
	          ),
	          cannedQueries
	        )
	      );
	    }
	  }]);

	  return GraphQLWebClientWrapper;
	})(_react2['default'].Component);

	exports['default'] = GraphQLWebClientWrapper;
	module.exports = exports['default'];

/***/ },
/* 3 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	var _get = function get(_x, _x2, _x3) { var _again = true; _function: while (_again) { var object = _x, property = _x2, receiver = _x3; desc = parent = getter = undefined; _again = false; var desc = Object.getOwnPropertyDescriptor(object, property); if (desc === undefined) { var parent = Object.getPrototypeOf(object); if (parent === null) { return undefined; } else { _x = parent; _x2 = property; _x3 = receiver; _again = true; continue _function; } } else if ('value' in desc) { return desc.value; } else { var getter = desc.get; if (getter === undefined) { return undefined; } return getter.call(receiver); } } };

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _GraphQLQueryInput = __webpack_require__(4);

	var _GraphQLQueryInput2 = _interopRequireDefault(_GraphQLQueryInput);

	var _GraphQLQueryResults = __webpack_require__(6);

	var _GraphQLQueryResults2 = _interopRequireDefault(_GraphQLQueryResults);

	var divStyle = {};

	var GraphQLWebClient = (function (_React$Component) {
	  function GraphQLWebClient(props) {
	    _classCallCheck(this, GraphQLWebClient);

	    _get(Object.getPrototypeOf(GraphQLWebClient.prototype), 'constructor', this).call(this, props);
	    this.state = {
	      query: this.props.defaultQuery,
	      response: '(no response received yet)'
	    };
	    this.queryEvent = null;
	    this.queryDelay = 100;
	  }

	  _inherits(GraphQLWebClient, _React$Component);

	  _createClass(GraphQLWebClient, [{
	    key: 'onInputChange',
	    value: function onInputChange(event) {
	      this.setState({ query: event.target.value });
	      window.location.hash = encodeURIComponent(this.state.query);
	      this.queryBackend();
	    }
	  }, {
	    key: 'componentDidMount',
	    value: function componentDidMount() {
	      this.queryBackend();
	    }
	  }, {
	    key: 'componentWillReceiveProps',
	    value: function componentWillReceiveProps(nextProps) {
	      this.state.query = nextProps.defaultQuery;
	      this.queryBackend();
	    }
	  }, {
	    key: 'queryBackend',
	    value: function queryBackend() {
	      var _this = this;

	      if (this.queryEvent !== null) {
	        clearTimeout(this.queryEvent);
	      }
	      this.queryEvent = setTimeout(function () {
	        var xhr = new XMLHttpRequest();
	        xhr.open('get', '' + _this.props.endpoint + '?q=' + _this.state.query, true);
	        xhr.onload = function () {
	          _this.setState({ response: xhr.responseText });
	        };
	        xhr.send();
	      }, this.queryDelay);
	    }
	  }, {
	    key: 'render',
	    value: function render() {
	      return _react2['default'].createElement(
	        'div',
	        { style: divStyle },
	        _react2['default'].createElement(_GraphQLQueryInput2['default'], { query: this.state.query, onChange: this.onInputChange.bind(this) }),
	        _react2['default'].createElement(_GraphQLQueryResults2['default'], { results: this.state.response })
	      );
	    }
	  }]);

	  return GraphQLWebClient;
	})(_react2['default'].Component);

	exports['default'] = GraphQLWebClient;
	module.exports = exports['default'];

/***/ },
/* 4 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _styles = __webpack_require__(5);

	var _styles2 = _interopRequireDefault(_styles);

	var GraphQLQueryInput = (function (_React$Component) {
	  function GraphQLQueryInput() {
	    _classCallCheck(this, GraphQLQueryInput);

	    if (_React$Component != null) {
	      _React$Component.apply(this, arguments);
	    }
	  }

	  _inherits(GraphQLQueryInput, _React$Component);

	  _createClass(GraphQLQueryInput, [{
	    key: 'render',
	    value: function render() {
	      return _react2['default'].createElement('textarea', { style: _styles2['default'].textareaStyle,
	        onChange: this.props.onChange,
	        value: this.props.query
	      });
	    }
	  }]);

	  return GraphQLQueryInput;
	})(_react2['default'].Component);

	exports['default'] = GraphQLQueryInput;
	module.exports = exports['default'];

/***/ },
/* 5 */
/***/ function(module, exports, __webpack_require__) {

	"use strict";

	Object.defineProperty(exports, "__esModule", {
	  value: true
	});
	exports["default"] = {
	  textareaStyle: {
	    width: "45%",
	    margin: "5px",
	    padding: "5px",
	    float: "left",
	    height: "400px",
	    fontFamily: "Consolas, monospace"
	  }
	};
	module.exports = exports["default"];

/***/ },
/* 6 */
/***/ function(module, exports, __webpack_require__) {

	'use strict';

	Object.defineProperty(exports, '__esModule', {
	  value: true
	});

	var _createClass = (function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ('value' in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; })();

	function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

	function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

	function _inherits(subClass, superClass) { if (typeof superClass !== 'function' && superClass !== null) { throw new TypeError('Super expression must either be null or a function, not ' + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) subClass.__proto__ = superClass; }

	var _react = __webpack_require__(1);

	var _react2 = _interopRequireDefault(_react);

	var _styles = __webpack_require__(5);

	var _styles2 = _interopRequireDefault(_styles);

	var GraphQLQueryResults = (function (_React$Component) {
	  function GraphQLQueryResults() {
	    _classCallCheck(this, GraphQLQueryResults);

	    if (_React$Component != null) {
	      _React$Component.apply(this, arguments);
	    }
	  }

	  _inherits(GraphQLQueryResults, _React$Component);

	  _createClass(GraphQLQueryResults, [{
	    key: 'render',
	    value: function render() {
	      return _react2['default'].createElement('textarea', { style: _styles2['default'].textareaStyle,
	        value: this.props.results,
	        defaultValue: 'no response recieved',
	        readOnly: true
	      });
	    }
	  }]);

	  return GraphQLQueryResults;
	})(_react2['default'].Component);

	exports['default'] = GraphQLQueryResults;
	module.exports = exports['default'];

/***/ }
/******/ ]);