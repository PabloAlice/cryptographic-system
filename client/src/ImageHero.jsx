import React from 'react'
import { Jumbotron, Image } from 'react-bootstrap'

const Placeholder = () => (
    <div className="text-center">
        <h1> Present Encryption algorithm </h1>
        <h3 className="mb-4">Let's encrypt that shit!</h3>
        <p className="text-muted"> Complete The inputs below and see the magic baby </p>
    </div>
)

const ImageHero = (props) => (
    <Jumbotron>
        {props.imageSrc ? <Image fluid src={props.imageSrc} /> : <Placeholder/>
        
        }
    </Jumbotron>
)

export default ImageHero