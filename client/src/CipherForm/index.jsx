import React from 'react'
import Form from './Form'
import {encrypt, decrypt} from '../api'

export default class CipherFormContainer extends React.Component {
    generateDataForm = () => {
        const formData = new FormData()
        Object.keys(this.state).forEach(key => {
            formData.append(key, this.state[key])
        })
        return formData
    }
    onEncrypt = async (e) => {
        e.preventDefault()
        console.log('encrypting', this.state)
        const formData = this.generateDataForm()
        const res = await encrypt(formData)
        console.log(res)

    }
    onDecrypt = async (e) => {
        e.preventDefault()
        console.log('decrypting', this.state)
        const formData = this.generateDataForm()
        const res = await decrypt(formData)
        console.log(res)
    }

    onChange = (e) => {
        const { name, value } = e.target
        this.setState({ [name]: value })
    }

    onFileChange = (e) => {
        const [file] = e.target.files
        this.setState({ file })
    }

    render() {
        return (
            <Form onChange={this.onChange} onEncrypt={this.onEncrypt} onDecrypt={this.onDecrypt} onFileChange={this.onFileChange}/>
        )
    }
}