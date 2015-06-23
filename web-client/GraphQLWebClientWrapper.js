import React from 'react';

import GraphQLWebClient from './GraphQLWebClient';
import ExtraHeaders from './ExtraHeaders';

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
      extraHeaders: [],
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
  onChangeHeaders(newHeaders) {
    this.setState({extraHeaders: newHeaders});
  }
  render() {
    var cannedQueries = this.state.cannedQueries.map((query) => {
      return (
          <li><a href="#" onClick={this.onCannedQueryClicked.bind(this)}>{query}</a></li>
      );
    });
    var buttonLabel = "Run Query";
    if (this.refs.client) {
      buttonLabel = ["Run Query", "Waiting for input", "Running Query"][this.state.queryState];
    }
    return (
      <div className="container-fluid">
      <strong>github.com/tmc/graphql - web-client</strong>
      <div className="form">
        <div>
          <label>endpoint:</label>
          <input size="50" defaultValue={this.state.endpoint} onChange={this.onChangeEndpoint.bind(this)} />
        </div>
        <div className="form-group">
          <button className="btn btn-primary" disabled={this.state.runAutomatically} onClick={this.onClickQuery.bind(this)}>{buttonLabel}</button>
          <label class="checkbox-inline">
            <input type="checkbox" onChange={this.onChangeRunAutomatically.bind(this)} defaultChecked={this.state.runAutomatically} />
            run automatically?
          </label>
          <label class="checkbox-inline">
            <input type="checkbox" onChange={this.onChangeShowParseResult.bind(this)} />
            just show parse result?
          </label>
        </div>
        <div className="form-group">
        <ExtraHeaders headers={this.state.extraHeaders} onChange={this.onChangeHeaders.bind(this)} />
        </div>
      </div>
        <hr/>
        <GraphQLWebClient
          ref="client"
          defaultQuery={this.state.defaultQuery}
          showParseResult={this.state.showParseResult}
          autoRun={this.state.runAutomatically}
          onQueryState={this.onChildQueryStateChange.bind(this)}
          endpoint={this.state.endpoint}
          extraHeaders={this.state.extraHeaders}
        />
        <ul style={styles.clear}>
          <li>Canned Queries:</li>
          {cannedQueries}
        </ul>
      </div>
    );
  }
}
