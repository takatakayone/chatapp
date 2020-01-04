import React from 'react';
import './ChatHistory.scss';
import Message from "../Message";

class ChatHistory extends React.Component{
    render() {
        const messages = this.props.chatHistory.map((msg, index) => <Message message={msg.data}/>);

        return(
            <div className="ChatHistory">
                <h2>Chat History</h2>
                {messages}
            </div>
        )
    }
}

export default ChatHistory;

