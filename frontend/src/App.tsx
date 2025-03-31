import { useState } from 'react'
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
        <div className="w-full h-full bg-gray-800 text-white grid grid-cols-[350px_1fr]">
            <div className="bg-gray-700"></div>
            <div></div>
        </div>
    )
}

export default App
