import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: JSON.parse(this.props.message)
    };
  }

  render() {
    let msg = this.state.message.user + ": " + this.state.message.body
    return <div className="Message">{msg}</div>;
  }
}

export default Message;
