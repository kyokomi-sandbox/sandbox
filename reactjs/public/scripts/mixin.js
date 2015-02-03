var Logger = {
  logging(str) {
    console.log(str);
  },
  componentWillMount() {
    this.logging("component will mount");
  },
  componentDidMount() {
    this.logging("component did mount");
  }
};

var Hello = React.createClass({
  mixins: [Logger],
  componentWillMount() {
    this.logging("Hello component will mount");
  },
  componentDidMount() {
    this.logging("Hello component did mount");
  },
  render() {
    this.logging("render");
    return <div>Hello</div>
  }
});

React.render(
  <Hello />,
  document.getElementById('mixin')
);
