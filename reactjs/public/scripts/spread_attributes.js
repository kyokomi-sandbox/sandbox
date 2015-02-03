var ImageText = React.createClass({
  render() {
    return (
      <span>{this.props.text}<img src={this.props.src} width={this.props.width} height={this.props.height} /></span>
    );
  }
});

var ImageText = React.createClass({
    render() {
      //var {text, ...other} = this.props;
      return (<span>{text}<img {...other} /></span>);
    }
});
