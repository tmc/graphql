import React from 'react';
import GraphQLQueryInput from './GraphQLQueryInput';
import GraphQLQueryResults from './GraphQLQueryResults';

var parentDivStyle = {
  display: "flex",
  flexWrap: "wrap"
};
var divStyle = {
  margin: "auto",
  width: "49%",
  border: "1px dotted #ccc"
};
var divStyleQuerying = { border: "1px solid #aaa" };
var divStylePreQuerying = { border: "1px solid #000" };
Object.assign(divStyleQuerying, divStyle);
Object.assign(divStylePreQuerying, divStyle);

var styles = [divStyle, divStylePreQuerying, divStyleQuerying];


export default class GraphQLWebClient extends React.Component {
  defaultProps: {
	extraheaders: []
  }
  constructor(props) {
    super(props);
    this.state = {
      query: this.props.defaultQuery,
      queryState: 0,
      response: `(no response received yet)`
    };
    this.queryEvent = null;
    this.queryDelay = 500;
    this.xhr = null;
  }
  onInputChange(value) {
    window.location.hash = value;
    this.setState({query: value});
    if (this.props.autoRun) {this.queryBackend(); }
  }
  componentDidMount() {
    this.queryBackend();
  }
  componentWillReceiveProps(nextProps) {
    if (nextProps.defaultQuery !== this.props.defaultQuery) {
      this.setState({query: nextProps.defaultQuery});
    }
    if (!this.props.autoRun && nextProps.autoRun) {this.queryBackend(); }
  }
  setQueryState(newState) {
    if (this.props.onQueryState) {
      this.props.onQueryState(newState);
    }
    this.setState({queryState: newState});
  }
  queryBackend() {
    var queryDelay = this.queryDelay;
    if (this.queryEvent !== null) { clearTimeout(this.queryEvent); }
    if (this.xhr !== null) {
      this.xhr.abort();
    } else {
      queryDelay = 0;
    }
    this.setQueryState(1);
    this.queryEvent = setTimeout(() => {
      this.setQueryState(2);
      var xhr = new XMLHttpRequest();
      xhr.open('get', `${this.props.endpoint}?q=${encodeURIComponent(this.state.query)}`, true);
      xhr.setRequestHeader('X-Trace-Id', '1');
      if (this.props.showParseResult) {
        xhr.setRequestHeader('X-GraphQL-Only-Parse', '1');
      }
      this.props.extraHeaders.forEach((h) => {
        xhr.setRequestHeader(h[0], h[1]);
      });
      xhr.onload = () => {
          this.setState({response: xhr.responseText});
          this.setQueryState(0);
      };
      xhr.send();
      this.xhr = xhr;
    }, queryDelay);
  }
  render() {
    return (
     <div style={parentDivStyle}>
       <div style={divStyle}>
       <GraphQLQueryInput query={this.state.query} onChange={this.onInputChange.bind(this)} />
       </div>
       <div style={styles[this.state.queryState]}>
       <GraphQLQueryResults results={this.state.response} />
       </div>
     </div>
    );
  }
}

