import React, { Component } from 'react';
import {Container} from 'react-bootstrap';
import ImageHero from './ImageHero';
import CipherForm from './CipherForm';

import 'bootstrap/dist/css/bootstrap.css';

class App extends Component {
  state = {
    encryptedSrc : null
  }

  onEncrypted = fileName => {
    this.setState({encryptedSrc: `static/${fileName}`})
  }

  render() {
    return (
      <div>
        <ImageHero src={this.state.encryptedSrc} />
        <Container>
          <CipherForm onEncrypted={this.onEncrypted}/>
        </Container>
      </div>
    );
  }
}

export default App;
