import React, { Component } from 'react';
import axios from 'axios';
import Button from "react-bootstrap/Button";

export default class Download extends Component {

    handleDownload(e) {
        e.preventDefault();

        axios({
            url: '/go/download',
            method: 'GET',
            responseType: 'blob', // important
        }).then((response) => {
            const url = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', 'file.txt');
            document.body.appendChild(link);
            link.click();
        });
    }

    render() {
        return (
            <div style={{marginTop: 50}}>
                <Button onClick={this.handleDownload}>Download file</Button>
            </div>
        )
    }
}
