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

var Box = React.createClass({
  getInitialState() {
    return {
      windowWidth: window.innerWidth
    };
  },
  handleResize(e) {
    this.setState({windowWidth: window.innerWidth});
  },
  componentDidMount() {
    window.addEventListener('resize', this.handleResize);
  },
  componentWillUnmount() {
    window.removeEventListener('resize', this.handleResize);
  },
  render() {
    return <div>Current window width: {this.state.windowWidth}</div>;
  }
});
React.render(<Box />, document.getElementById('counter'));