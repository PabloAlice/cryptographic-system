import React, { Component } from 'react';
import {Container} from 'react-bootstrap';
import ImageHero from './ImageHero';
import CipherForm from './CipherForm';

import 'bootstrap/dist/css/bootstrap.css';

class App extends Component {
  render() {
    return (
      <div>
        <ImageHero />
        <Container>
          <CipherForm/>
        </Container>
      </div>
    );
  }
}

export default App;
