// 対応状況の一括変更モーダル
var Collectively = React.createClass({
  getInitialState: function() {
    return {
      selectOpt: "active"
    };
  },
  handleChange: function(event) {
    console.log("Collectively:handleChange: ", event.target.value);

    this.setState({selectOpt: event.target.value});
  },
  onClickSave: function() {
    console.log("Collectively:", this.state.selectOpt);
    this.props.onChange(this.state.selectOpt);
  },
  render: function() {
    return (
      <section role="collectively" className="Modal__change-status">
        <h1 className="heading-1">変更する</h1>
        <div className="status-changer">
          <label className="label">ステータス</label>
          <div className="custom-selector">
            <select onChange={this.handleChange}>
              <option value="active">対応中</option>
              <option value="started">入力中</option>
              <option value="complete">回答済</option>
              <option value="closed">クローズ</option>
              <option value="archived">アーカイヴ</option>
            </select>
          </div>
        </div>
        <div className="changer settlement">
          <button className="save closer" onClick={this.onClickSave}><span>変更</span></button>
        </div>
      </section>
    );
  }
});

// ListHeader is Header for ThreadList
var ListHeader = React.createClass({
  getInitialState: function() {
    return {
      modalFlag: false,
      timeout: null,
      modalClassName: ""
    };
  },
  // 一括ステータス更新のModal表示
  modalOpen: function() {
    console.log("ListHeader:modalOpen: ");

    this.setState({modalClassName: "allselect-live"});

    var self = this;
    this.state.timeout = setTimeout(function() {
      console.log("ListHeader:setTimeout");

      self.setState({modalClassName: "allselect-live allselect-transit"});
      self.state.timeout = null;
    }, 1000);
  },
  // 一括ステータス更新のModal閉じる
  modalClose: function() {
    console.log("ListHeader:modalClose: ");

    this.setState({modalClassName: "allselect-live"});

    var self = this;
    this.state.timeout = setTimeout(function() {
      console.log("ListHeader:setTimeout");

      self.setState({modalClassName: "allselect-alive"});
      self.state.timeout = null;
    }, 1000);

    // TODO: すべてのチェックをはずす
    //$allChatSelectChackModal.closeCallback =->
    //$chatCheckBoxies.removeAttr 'checked'
  },
  // 全選択のチェックボックス変更時
  onChange: function() {
    console.log("ListHeader:onChange: ");

    if (this.state.modalFlag) {
      this.modalClose();
      this.setState({modalFlag: false});
    } else {
      this.modalOpen();
      this.setState({modalFlag: true});
    }
  },
  // モーダル内でステータス変更時
  onChangeCollectively: function(selectOpt) {
    console.log("ListHeader:onChangeCollectively: ", selectOpt);
  },
  render: function() {
    return (
      <header className="chat-room-list-header main-list-header">
        <div className={"modals " + this.state.modalClassName}>
          <Collectively onChange={this.onChangeCollectively} />s
        </div>
        <div className="allselecter">
          <input id="allselect" type="checkbox" onChange={this.onChange} />
          <label htmlFor="allselect" className="label">全て選択</label>
        </div>
      </header>
    );
  }
});

React.render(
  <ListHeader />,
  document.getElementById('listHeader')
);
