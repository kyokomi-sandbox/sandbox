var Counter = React.createClass({
  getInitialState() {
    return {
      count: 0
    };
  },
  onClick() {
    this.setState({count: this.state.count + 1});
  },
  render() {
    return (
      <div>
        <div>count:{this.state.count}</div>
        <button onClick={this.onClick}>click!</button>
      </div>
    );
  }
});

React.render(
  <Counter />,
  document.getElementById('counter')
);