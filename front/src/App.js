import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import ChatHistory from './components/ChatHistory/ChatHistory';
import ChatInput from './components/ChatInput/ChatInput';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: []
    }
  }
  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory: [...prevState.chatHistory, msg]
      }))
      console.log(this.state);
    });
  }
  send(event) {
    if(event.keyCode === 13 && event.target.value.length > 0) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }
  render() {
    var ph = "Type your name and press Enter to login"
    if (this.state.chatHistory.length > 0) {
      ph = "Type message and press Enter"
    }
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} placeholder={ph} />
      </div>
    );
  }
}

export default App;
