import { useEffect, useState } from 'react'
import { Greet, ListDbTables, ListTables } from '../wailsjs/go/app/App'
import DbListItem from './components/DbListItem'

function App() {
    const [dbList, setDbList] = useState<string[]>([])
    const [tableList, setTableList] = useState<string[]>([])
    const [selectedDb, setSelectedDb] = useState<string>('')
    const [resultText, setResultText] = useState('Please enter your name below 👇')
    const [name, setName] = useState('')
    const updateName = (e: any) => setName(e.target.value)
    const updateResultText = (result: string) => setResultText(result)

    useEffect(() => {
        if (selectedDb == '') {
            ListDbTables().then((res) => {
                console.log(res)
                if (res == null) {
                    setDbList([])
                    return
                }
                setDbList(res)
            })
        } else {
            ListTables(selectedDb).then((res) => {
                console.log(res)
                if (res == null) {
                    setTableList([])
                    return
                }
                setTableList(res)
            })
        }
    }, [selectedDb])

    function clickDbListItem(name: string) {
        setSelectedDb(name)
    }

    function backButton() {
        setSelectedDb('')
    }

    return (
        <div className="w-full h-full bg-gray-800 text-white grid grid-cols-[350px_1fr]">
            <div className="bg-gray-700 py-3 max-h-full overflow-y-auto">
                {selectedDb != '' ? (
                    <div>
                        <div className="mb-3 text-center text-xl">
                            <button
                                className="cursor-pointer py-1 px-3 me-3"
                                onClick={() => {
                                    backButton()
                                }}
                            >
                                {'<'}
                            </button>
                            Database: {selectedDb}
                        </div>
                        <div className="flex flex-col gap-3 px-5">
                            {tableList.map((table) => (
                                <DbListItem
                                    key={table}
                                    title={table}
                                    onClick={() => {
                                        console.log('table ' + table + ' clicked')
                                    }}
                                />
                            ))}
                        </div>
                    </div>
                ) : (
                    <div className="flex flex-col gap-3 px-5">
                        {dbList.map((dbListItem) => (
                            <DbListItem
                                key={dbListItem}
                                title={dbListItem}
                                onClick={() => {
                                    clickDbListItem(dbListItem)
                                }}
                            />
                        ))}
                    </div>
                )}
            </div>
            <div></div>
        </div>
    )
}

export default App
