import { useEffect } from 'react'
import { GetDatabaseConfigs } from '../../wailsjs/go/app/App'
import InputText from '../components/form/InputText'
import './../style.css'

export default function Dashboard() {
    useEffect(() => {
        GetDatabaseConfigs()
            .then((result) => {
                console.log(result)
            })
            .catch((error) => {
                alert(error)
            })
    }, [])

    return (
        <div className="h-full w-full">
            <div className="w-full">
                <h1>Create Connection</h1>
                <div className="grid grid-cols-2 gap-5">
                    <div className="grid grid-cols-[auto_1fr] gap-3">
                        <label>Host:</label>
                        <InputText />
                    </div>
                    <div className="grid grid-cols-[auto_1fr] gap-3">
                        <label>Port:</label>
                        <InputText />
                    </div>
                </div>
            </div>
            <div>Available Connections</div>
        </div>
    )
}
