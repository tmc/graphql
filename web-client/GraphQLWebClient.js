import React from 'react';
import GraphQLQueryInput from './GraphQLQueryInput';
import GraphQLQueryResults from './GraphQLQueryResults';

var divStyle = {
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
  onInputChange(event) {
    this.setState({query: event.target.value});
    window.location.hash = encodeURIComponent(this.state.query);
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
      xhr.onload = () => {
          this.setState({response: xhr.responseText});
      };
      xhr.send();
    }, this.queryDelay);
  }
  render() {
    return (
     <div style={divStyle}>
       <GraphQLQueryInput query={this.state.query} onChange={this.onInputChange.bind(this)} />
       <GraphQLQueryResults results={this.state.response} />
     </div>
    );
  }
}

