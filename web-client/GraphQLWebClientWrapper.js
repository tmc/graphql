import React from 'react';

import GraphQLWebClient from './GraphQLWebClient';

var styles = {
    clear: {clear: "both"}
};

export default class GraphQLWebClientWrapper extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      endpoint: this.props.endpoint,
      cannedQueries: [
        `{ __schema { root_fields { name, description } } }`,
        `{ __types { name, description} }`,
        `{ __types { name, description, fields { name, description } } }`,
        `{ TodoUserClass{ objectId, name, lists:TodoItemListClass_owner { objectId, name, items:TodoItemClass_list { objectId, done, description } } } }`
      ]
    };
    this.state.defaultQuery = this.state.cannedQueries[0];
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
