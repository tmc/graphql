import React from 'react';
import AceEditor from 'react-ace';

var brace = require('brace');
require('brace/mode/javascript');
require('brace/theme/github');

export default class GraphQLQueryInput extends React.Component {
  render() {
    return (
      <AceEditor
          theme="github"
          value={this.props.query}
          name="input"
          width="100%"
          wordWrap={true}
          onChange={this.props.onChange} />
    );
  }
}
