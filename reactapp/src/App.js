import React, { Component } from 'react';
import { Card, Button, Container}  from "react-bootstrap";

import '../node_modules/bootstrap/dist/css/bootstrap.min.css';

class App extends Component {
    render() {
        return (
            <Container>
                <Card border="dark" style={{ width: '18rem'}}>
                    <Card.Img src="/logo512.png"/>
                    <Card.Body>
                        <Card.Title>
                            React & Golang file download
                        </Card.Title>
                        <Card.Text>
                            XHR approach requires more effort.
                            That is why lets stick with HTML.
                        </Card.Text>
                    </Card.Body>
                    <Card.Footer>
                        <form enctype="multipart/form-data" action="/go/download" method="post">
                            <Button id="image-file" type="file">Download file</Button>
                        </form>
                    </Card.Footer>
                </Card>
            </Container>
        );
    }
}

export default App;
