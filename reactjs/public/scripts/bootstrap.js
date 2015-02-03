/**
 * Created by kyokomi on 15/02/04.
 */
var Jumbotron = ReactBootstrap.Jumbotron;
var Col = ReactBootstrap.Col;

var HelloBoot = React.createClass({
  render: function() {
    return (
      <Col xs={6} md={4}>
        <Jumbotron>
          <p>{this.props.title}</p>
        </Jumbotron>
      </Col>
    )
  }
});

React.render(
  <HelloBoot title="hoge" />,
  document.getElementById('boot')
);
