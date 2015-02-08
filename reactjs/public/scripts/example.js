/**
 * This file provided by Facebook is for non-commercial testing and evaluation purposes only.
 * Facebook reserves all rights not expressly granted.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * FACEBOOK BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

var React = require("React");
var $ = require("jquery");

var converter = new Showdown.converter();

var Comment = React.createClass({
  render: function() {
    var rawMarkup = converter.makeHtml(this.props.children.toString());
    //var rawMarkup = this.props.children.toString();
    return (
      <div className="comment">
        <h2 className="commentAuthor">
          {this.props.author}
        </h2>
        <span dangerouslySetInnerHTML={{__html: rawMarkup}} />
      </div>
    );
  }
});

var CommentBox = React.createClass({
  loadCommentsFromServer: function() {
    $.ajax({
      url: this.props.url,
      dataType: 'json',
      success: function(data) {
        this.setState({data: data});
      }.bind(this),
      error: function(xhr, status, err) {
        console.error(this.props.url, status, err.toString());
      }.bind(this)
    });
  },
  handleCommentSubmit: function(comment) {
    var comments = this.state.data;
    comments.push(comment);
    this.setState({data: comments}, function() {
      // `setState` accepts a callback. To avoid (improbable) race condition,
      // `we'll send the ajax request right after we optimistically set the new
      // `state.
      $.ajax({
        url: this.props.url,
        dataType: 'json',
        type: 'POST',
        data: comment,
        success: function(data) {
          this.setState({data: data});
        }.bind(this),
        error: function(xhr, status, err) {
          console.error(this.props.url, status, err.toString());
        }.bind(this)
      });
    });
  },
  getInitialState: function() {
    return {data: []};
  },
  componentDidMount: function() {
    this.loadCommentsFromServer();
    setInterval(this.loadCommentsFromServer, this.props.pollInterval);
  },
  render: function() {
    return (
      <div className="commentBox">
        <h1>Comments</h1>
        <CommentList data={this.state.data} />
        <CommentForm onCommentSubmit={this.handleCommentSubmit} />
      </div>
    );
  }
});

var CommentList = React.createClass({
  render: function() {
    var commentNodes = this.props.data.map(function(comment, index) {
      return (
        // `key` is a React-specific concept and is not mandatory for the
        // purpose of this tutorial. if you're curious, see more here:
        // http://facebook.github.io/react/docs/multiple-components.html#dynamic-children
        <Comment author={comment.author} key={index}>
          {comment.text}
        </Comment>
      );
    });
    return (
      <div className="commentList">
        {commentNodes}
      </div>
    );
  }
});

var CommentForm = React.createClass({
  handleSubmit: function(e) {
    e.preventDefault();
    var author = this.refs.author.getDOMNode().value.trim();
    var text = this.refs.text.getDOMNode().value.trim();
    if (!text || !author) {
      return;
    }
    this.props.onCommentSubmit({author: author, text: text});
    this.refs.author.getDOMNode().value = '';
    this.refs.text.getDOMNode().value = '';
    return;
  },
  render: function() {
    return (
      <form className="commentForm" onSubmit={this.handleSubmit}>
        <input type="text" placeholder="Your name" ref="author" />
        <input type="text" placeholder="Say something..." ref="text" />
        <input type="submit" value="Post" />
      </form>
    );
  }
});

//React.render(
//  <CommentBox url="comments.json" pollInterval={2000} />,
//  document.getElementById('content')
//);

//var Text = React.createClass({
//  getInitialState() {
//    return {
//      textValue: "initial value"
//    };
//  },
//  changeText1(e) {
//    this.setState({textValue: e.target.value});
//  },
//  changeText2(e) {
//    this.setState({textValue: this.refs.inputText.getDOMNode().value});
//  },
//  render() {
//    return (
//      <div>
//        <p>{this.state.textValue}</p>
//        <input type="text" value={this.state.textValue} onChange={this.changeText1} />
//        <input type="text" ref="inputText" defaultValue="initial value" />
//        <button onClick={this.changeText2}>change</button>
//      </div>
//    );
//  }
//});

//React.render(
//  <Text />,
//  document.getElementById('textForm')
//);

//var OreTextArea = React.createClass({
//  getInitialState() {
//    return {
//      textAreaValue: "initial value"
//    };
//  },
//  onChangeText(e) {
//    this.setState({textAreaValue: e.target.value});
//  },
//  onClick() {
//    this.setState({textAreaValue: this.refs.textArea.getDOMNode().value});
//  },
//  render() {
//    return (
//      <div>
//        <div>{this.state.textAreaValue}</div>
//        <div>
//          <textarea value={this.state.textAreaValue} onChange={this.onChangeText} />
//        </div>
//        <div>
//          <textarea ref="textArea">this is default value</textarea>
//          <button onClick={this.onClick}>change</button>
//        </div>
//      </div>
//    );
//  }
//});

//React.render(
//  <OreTextArea />,
//  document.getElementById('oreTextArea')
//);

//var OreSelectBox = React.createClass({
//  getDefaultProps() {
//    return {
//      answers: [1, 10, 100, 1000]
//    };
//  },
//  getInitialState() {
//    return {
//      selectValue: 1,
//      selectValues: [1,100]
//    };
//  },
//  onChangeSelectValue(e) {
//    this.setState({selectValue: e.target.value});
//  },
//  onChangeSelectValues(e) {
//    var idx = 0;
//  	var values = [];
//  	for (var i = 0; i < e.target.options.length; i++) {
//  	    var opt = e.target.options[i]
//  	    if (!opt.selected) {
//            continue;
//        }
//        values[idx] = opt.value;
//        idx++;
//  	}
//
//    this.setState({selectValues: values});
//  },
//  render() {
//    var options = this.props.answers.map(function(answer) {
//      return <option value={answer} key={answer}>{answer}</option>;
//    });
//    return (
//      <div>
//        <div>selectValue: {this.state.selectValue}</div>
//        <div>
//          <select value={this.state.selectValue} onChange={this.onChangeSelectValue}>
//            {options}
//          </select>
//        </div>
//        <div>selectValues: {this.state.selectValues.join(",")}</div>
//        <div>
//          <select multiple={true} defaultValue={this.state.selectValues} onChange={this.onChangeSelectValues}>
//            {options}
//          </select>
//        </div>
//      </div>
//    );
//  }
//});

//React.render(
//  <OreSelectBox />,
//  document.getElementById('oreSelectBox')
//);

module.exports = CommentBox;

