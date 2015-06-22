import React from 'react';
import AceEditor from 'react-ace';

var brace = require('brace');
require('brace/mode/javascript');
require('brace/theme/github');

export default class GraphQLQueryResults extends React.Component {
  render() {
    return (
      <AceEditor
          mode="javascript"
          theme="github"
          showPrintMargin={false}
          showGutter={false}
          width="100%"
          value={this.props.results}
          defaultValue={"no response recieved"}
          name="results"
          readOnly={true}
      />
    );
  }
}
