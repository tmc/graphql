import React from 'react';
import AceEditor from 'react-ace';

import styles from './styles';

var brace = require('brace');
require('brace/mode/javascript');
require('brace/theme/github');

export default class GraphQLQueryInput extends React.Component {
  render() {
    return (
      <AceEditor
          mode="javascript"
          theme="github"
          showGutter={false}
          value={this.props.query}
          name="input"
          wordWrap={true}
          onChange={this.props.onChange} />
    );
  }
}
