const URL = "http://localhost:3000/api"
export async function encrypt(body) {
    const response = await fetch(`${URL}/encryption`, {
        method: 'POST',
        body
    })
    return await response.json()
}

export async function decrypt(body) {
    const response = await fetch(`${URL}/decryption`, {
        method: 'POST',
        body
    })
    return response.json()
}