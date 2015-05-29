import React from 'react';

import GraphQLWebClient from './GraphQLWebClient';

var styles = {
    clear: {clear: "both"}
};

export default class GraphQLWebClientWrapper extends React.Component {
  constructor(props) {
    super(props);
    var endpoint = this.props.endpoint;
    if (window && window.location.search) {
      var m = window.location.search.match(/endpoint=(.+)/)[1];
      if (m) { endpoint = m; }
    }
    this.state = {
      endpoint: endpoint,
      cannedQueries: [
        `{ __schema { root_fields { name, description } } }`,
        `{ __types { name, description} }`,
        `{ __types { name, description, fields { name, description } } }`,
        `{ _User { __type__ { fields { name } } } }`
      ]
    };
    this.state.defaultQuery = this.state.cannedQueries[0];
    if (window.location.hash.length > 1) {
      this.state.defaultQuery = decodeURIComponent(window.location.hash.slice(1));
    }
  }
  onChange(event) {
    this.setState({endpoint: event.target.value});
  }
  onCannedQueryClicked(event) {
    this.setState({defaultQuery: event.target.text});
  }
  render() {
    var cannedQueries = this.state.cannedQueries.map((query) => {
      return (
          <li><a href="#" onClick={this.onCannedQueryClicked.bind(this)}>{query}</a></li>
      );
    });
    return (
      <div>
        <h1>graphql client</h1>
        <label>graphql endpoint:</label>
        <input size="50" defaultValue={this.state.endpoint} onChange={this.onChange.bind(this)} />
        <hr/>
        <GraphQLWebClient
          defaultQuery={this.state.defaultQuery}
          endpoint={this.state.endpoint}
        />
        <ul style={styles.clear}>
          <li>Canned Queries:</li>
          {cannedQueries}
        </ul>
      </div>
    );
  }
}
