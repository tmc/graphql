import React from 'react';

import styles from './styles';

export default class GraphQLQueryInput extends React.Component {
  render() {
    return (
      <textarea style={styles.textareaStyle}
        onChange={this.props.onChange}
        value={this.props.query}
      />
    );
  }
}
