var Hello2 = React.createClass({
  getDefaultProps() {
    return {
      tag: "React"
    };
  },
  // DOM追加時は呼ばれない。このDOMがsetPropsされたり、上書きされるときに呼ばれる
  componentWillReceiveProps(nextProps) {
    console.log("componentWillReceiveProps !!!!", nextProps);
  },
  componentWillUpdate(nextProps, nextState) {
    console.log("componentWillUpdate !!!!", nextProps, nextState);
  },
  render() {
    return <div>{this.props.tag}:Hello {this.props.name} => {this.props.children}</div>
  }
});

var component = React.render(
	<Hello2 name="World" />,
	document.getElementById('hello2')
);
// <div>React:Hello World</div>

// setPropsするので、componentWillReceivePropsが呼ばれる
component.setProps({ name: "foo" });      // <div>React:Hello foo</div>
// replacePropsするので、componentWillReceivePropsが呼ばれる
component.replaceProps({ name: "hoge" }); // <div>:Hello hoge</div>

var idx = 0;

[
  <Hello2>xxx</Hello2>,
  <Hello2><span>1</span><span>2</span></Hello2>,
  <Hello2></Hello2>
].forEach( jsx => {
  idx++;
  // ループ2回目のやつがhello2なので、前のComponentが書き換えられて、componentWillReceivePropsが呼ばれる
  var children = React.render(jsx, document.getElementById('hello' + idx)).props.children;
  console.log("#########" + children + "##########");
  console.log(React.Children.count(children));
  React.Children.forEach(children, (child) => { console.log(child) });
});