import React from 'react';

import styles from './styles';

export default class GraphQLQueryResults extends React.Component {
  render() {
    return (
      <textarea style={styles.textareaStyle}
        value={this.props.results}
        defaultValue={"no response recieved"}
		readOnly={true}
	  />
    );
  }
}
