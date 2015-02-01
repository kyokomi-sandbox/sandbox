var Hello2 = React.createClass({
  getDefaultProps() {
    return {
      tag: "React"
    };
  },
  render() {
    return <div>{this.props.tag}:Hello {this.props.name}</div>
  }
});

var component = React.render(
	<Hello2 name="World" />,
	document.getElementById('hello2')
);
// <div>React:Hello World</div>

component.setProps({ name: "foo" });      // <div>React:Hello foo</div>
component.replaceProps({ name: "hoge" }); // <div>:Hello hoge</div>