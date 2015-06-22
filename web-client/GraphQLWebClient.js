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

export default class GraphQLWebClient extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      query: this.props.defaultQuery,
      response: `(no response received yet)`
    };
    this.queryEvent = null;
    this.queryDelay = 100;
  }
  onInputChange(value) {
    window.location.hash = value;
    this.setState({query: value});
    this.queryBackend();
  }
  componentDidMount() {
    this.queryBackend();
  }
  componentWillReceiveProps(nextProps) {
    this.state.query = nextProps.defaultQuery;
    this.queryBackend();
  }
  queryBackend() {
    if (this.queryEvent !== null) { clearTimeout(this.queryEvent); }
    this.queryEvent = setTimeout(() => {
      var xhr = new XMLHttpRequest();
      xhr.open('get', `${this.props.endpoint}?q=${this.state.query}`, true);
      xhr.setRequestHeader('X-Trace-Id', '1');
      if (this.props.showParseResult) {
        xhr.setRequestHeader('X-GraphQL-Only-Parse', '1');
      }
      xhr.onload = () => {
          this.setState({response: xhr.responseText});
      };
      xhr.send();
    }, this.queryDelay);
  }
  render() {
    return (
     <div style={parentDivStyle}>
       <div style={divStyle}>
       <GraphQLQueryInput query={this.state.query} onChange={this.onInputChange.bind(this)} />
       </div>
       <div style={divStyle}>
       <GraphQLQueryResults results={this.state.response} />
       </div>
     </div>
    );
  }
}

