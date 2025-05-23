import { useEffect, useState } from 'react'
import { ListDbTables, ListTables } from '../wailsjs/go/app/App'
import { EventsOn, EventsOff } from '../wailsjs/runtime/runtime'
import DbListItem from './components/DbListItem'
import Dashboard from './pages/Dashboard'
import DbOverview from './pages/DbOverview'
import './style.css'
import TableDataview from './pages/TableDataview'

enum Pages {
    Dashboard = 'dashboard',
    DbOverview = 'dbOverview',
    TableDataview = 'tableDataview',
}

function App() {
    const [selectedPage, setSelectedPage] = useState<Pages>(Pages.Dashboard)
    const [dbList, setDbList] = useState<string[]>([])
    const [tableList, setTableList] = useState<string[]>([])
    const [selectedDb, setSelectedDb] = useState<string>('')
    const [selectedTable, setSelectedTable] = useState<string>('')
    // const [name, setName] = useState('')
    // const updateName = (e: any) => setName(e.target.value)

    useEffect(() => {
        const handler = () => {
            console.log('connection changed!')
            updateMenuData()
        }

        EventsOn('ConnectionChanged', handler)

        return () => {
            EventsOff('ConnectionChanged')
        }
    }, [])

    useEffect(() => {
        updateMenuData()
    }, [selectedDb])

    function updateMenuData() {
        if (selectedDb == '') {
            ListDbTables()
                .then((res) => {
                    console.log(res)
                    if (res == null) {
                        setDbList([])
                        return
                    }
                    setDbList(res)
                })
                .catch((error) => {
                    setDbList([])
                })
        } else {
            ListTables(selectedDb)
                .then((res) => {
                    console.log(res)
                    if (res == null) {
                        setTableList([])
                        return
                    }
                    setTableList(res)
                })
                .catch((error) => {
                    setDbList([])
                })
        }
    }

    function clickTableListItem(name: string) {
        setSelectedTable(name)
        setSelectedPage(Pages.TableDataview)
    }

    function clickDbListItem(name: string) {
        setSelectedDb(name)
        setSelectedPage(Pages.DbOverview)
    }

    function backButton() {
        setSelectedDb('')
        setSelectedTable('')
        setSelectedPage(Pages.Dashboard)
    }

    function dashboardButton() {
        setSelectedPage(Pages.Dashboard)
        setSelectedDb('')
        setSelectedTable('')
    }

    return (
        <div className="w-full h-full bg-gray-800 text-white grid grid-cols-[350px_1fr]">
            <div className="bg-gray-700 py-3 max-h-full overflow-y-auto">
                <div className="border-b border-gray-500 pb-3 mb-3">
                    <div className="flex flex-col gap-3 px-5">
                        <DbListItem
                            title="Dashboard"
                            onClick={() => {
                                dashboardButton()
                            }}
                        />
                    </div>
                </div>
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
                                    selected={selectedTable == table}
                                    onClick={() => {
                                        clickTableListItem(table)
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
            <div className="p-3">
                {selectedPage == Pages.Dashboard ? (
                    <Dashboard />
                ) : selectedPage == Pages.DbOverview ? (
                    <DbOverview />
                ) : selectedPage == Pages.TableDataview ? (
                    <TableDataview />
                ) : null}
            </div>
        </div>
    )
}

export default App
