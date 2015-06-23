import React from 'react';

export default class ExtraHeaders extends React.Component {
  getDefaultProps: {
    headers: []
  }
  renderHeader(h, idx) {
    return (
      <div key={idx}>
        <label>name:</label><input size={30} defaultValue={h[0]} onChange={(e) => {this.onChange(e, idx, 0); }} />
        <label>value:</label><input size={30} defaultValue={h[1]} onChange={(e) => {this.onChange(e, idx, 1); }} />
      </div>
    );
  }
  onChange(e, i, j) {
    var hdrs = this.props.headers.slice();
    hdrs[i][j] = e.target.value;
    this.props.onChange(hdrs);
  }
  onClick() {
    var hdrs = this.props.headers.slice();
    hdrs.push(["X-", ""]);
    this.props.onChange(hdrs);
  }
  onClickPop() {
    var hdrs = this.props.headers.slice();
    hdrs.pop();
    this.props.onChange(hdrs);
  }
  render() {
    return (
      <div>
        <button className="btn btn-xs" onClick={this.onClick.bind(this)}>add extra header</button>
        <button className="btn btn-danger btn-xs" onClick={this.onClickPop.bind(this)}>pop extra header</button>
        {this.props.headers.map(this.renderHeader.bind(this))}
      </div>
    );
  }
}
