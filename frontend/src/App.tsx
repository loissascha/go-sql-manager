import { useState } from 'react'
import logo from './assets/images/logo-universal.png'
import './App.css'
import { Greet, ListDbTables } from '../wailsjs/go/app/App'

function App() {
    const [resultText, setResultText] = useState('Please enter your name below 👇')
    const [name, setName] = useState('')
    const updateName = (e: any) => setName(e.target.value)
    const updateResultText = (result: string) => setResultText(result)

    function greet() {
        Greet(name).then(updateResultText)
    }

    function listDbButton() {
        ListDbTables().then((res) => {
            console.log(res)
        })
    }

    return (
        <div id="App" className='w-full h-full'>
            <img src={logo} id="logo" alt="logo" />
            <div id="result" className="result">
                {resultText}
            </div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>
                    Greet
                </button>
                <button
                    onClick={() => {
                        listDbButton()
                    }}
                >
                    List Databases
                </button>
            </div>
        </div>
    )
}

export default App
