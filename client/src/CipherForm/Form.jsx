import React from 'react'
import { Form, Col, Button, Row } from 'react-bootstrap'

const CipherForm = (props) => (
    <Form onSubmit={props.onEncrypt}>
        <Form.Row>
            <Form.Group as={Col}>
                <Form.Label className="h5">File to Encrypt</Form.Label>
                <Form.Control type="file" size="lg" name="file" onChange={props.onFileChange} required/>
            </Form.Group>
            <Form.Group as={Row}>
                <Form.Label className="h5">
                    Chaining Methods
                </Form.Label>
                <Col>
                    <Form.Check
                        onChange={props.onChange}
                        type="radio"
                        label="ECB"
                        value="ECB"
                        name="method"
                        required
                    />
                </Col>
                <Col>
                    <Form.Check
                        onChange={props.onChange}
                        type="radio"
                        label="CBC"
                        value="CBC"
                        name="method"
                        required
                    />
                </Col>
                <Col>
                    <Form.Check
                        onChange={props.onChange}
                        type="radio"
                        label="CFB"
                        value="CFB"
                        name="method"
                        required
                    />
                </Col>
                <Col>
                    <Form.Check
                        onChange={props.onChange}
                        type="radio"
                        label="OFB"
                        value="OFB"
                        name="method"
                        required
                    />
                </Col>
            </Form.Group>
        </Form.Row>

        <Form.Group>
            <Form.Label className="h5">Key</Form.Label>
            <Form.Control placeholder="encryption Key" size="lg" onChange={props.onChange} name="key" required/>
        </Form.Group>

        <Form.Group>
            <Form.Label className="h5">IV</Form.Label>
            <Form.Control placeholder="inicialization value" size="lg" onChange={props.onChange} name="iv" required/>
        </Form.Group>

        <Form.Row>
            <Col>
                <Button variant="primary" type="submit" name="encrypt" block size="lg">
                    Encrypt
                </Button>
            </Col>
            <Col>
                <Button variant="primary" type="button" name="decrypt" block size="lg" onClick={props.onDecrypt}>
                    Decrypt
                </Button>
            </Col>
        </Form.Row>
    </Form>
)

export default CipherForm