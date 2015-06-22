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
      runAutomatically: true,
      queryState: 0,
      showParseResult: false,
      cannedQueries: [
        `{ __schema { root_fields { name, description } } }`,
        `{ __schema { types { name, description } } }`,
        `{ __schema { types { name, description, fields { name, description } } } }`,
        `{ _User { __type__ { fields { name } } } }`,
        `{ _User { objectId } }`
      ]
    };
    this.state.defaultQuery = this.state.cannedQueries[0];
    if (window.location.hash.length > 1) {
      this.state.defaultQuery = decodeURIComponent(window.location.hash.slice(1));
    }
  }
  onChangeEndpoint(event) {
    this.setState({endpoint: event.target.value});
  }
  onChangeShowParseResult(event) {
    this.setState({showParseResult: event.target.checked});
  }
  onChangeRunAutomatically(event) {
    console.log("ocra:", event.target.checked);
    this.setState({runAutomatically: event.target.checked});
  }
  onCannedQueryClicked(event) {
    this.setState({defaultQuery: event.target.text});
  }
  onChildQueryStateChange(newState) {
    this.setState({queryState: newState});
  }
  onClickQuery(state) {
    this.refs.client.queryBackend();
  }
  render() {
    var cannedQueries = this.state.cannedQueries.map((query) => {
      return (
          <li><a href="#" onClick={this.onCannedQueryClicked.bind(this)}>{query}</a></li>
      );
    });
    var buttonLabel = "Run Query";
    if (this.refs.client) {
      console.log(this.refs.client);
      buttonLabel = ["Run Query", "Waiting for input", "Running Query"][this.state.queryState];
      console.log(buttonLabel);
    }
    return (
      <div>
        <strong>github.com/tmc/graphql - web-client</strong>
        <p>
        <label htmlFor="_graphql_endpoint">endpoint:</label>
        <input id="_graphql_endpoint" size="50" defaultValue={this.state.endpoint} onChange={this.onChangeEndpoint.bind(this)} />
        </p>
        <p>
        <label htmlFor="_graphql_run_auto">run automatically? </label>
        <input id="_graphql_run_auto" type="checkbox" onChange={this.onChangeRunAutomatically.bind(this)} defaultChecked={this.state.runAutomatically} />
        <span> </span>
        <label htmlFor="_graphql_show_parse">just show parse result?</label>
        <input id="_graphql_show_parse" type="checkbox" onChange={this.onChangeShowParseResult.bind(this)} />
        <span> </span>
        <button className="btn btn-primary" disabled={this.state.runAutomatically} onClick={this.onClickQuery.bind(this)}>{buttonLabel}</button>
        </p>
        <hr/>
        <GraphQLWebClient
          ref="client"
          defaultQuery={this.state.defaultQuery}
          showParseResult={this.state.showParseResult}
          autoRun={this.state.runAutomatically}
          onQueryState={this.onChildQueryStateChange.bind(this)}
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
