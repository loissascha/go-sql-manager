import { useState, useEffect } from "react";
import { GreetService } from "../bindings/github.com/loissascha/go-sql-manager/internal/services";
import { Events, WML } from "@wailsio/runtime";

function App() {
    const [name, setName] = useState("");
    const [result, setResult] = useState("Please enter your name below 👇");
    const [time, setTime] = useState("Listening for Time event...");

    const doGreet = () => {
        let localName = name;
        if (!localName) {
            localName = "anonymous";
        }
        GreetService.Greet(localName)
            .then((resultValue) => {
                setResult(resultValue);
            })
            .catch((err) => {
                console.log(err);
            });
    };

    useEffect(() => {
        Events.On("time", (timeValue) => {
            setTime(timeValue.data);
        });
        // Reload WML so it picks up the wml tags
        WML.Reload();
    }, []);

    return (
        <div className="text-2xl w-full h-full bg-gray-100 dark:bg-gray-900 dark:text-white">
            Hello w
        </div>
    );
}

export default App;
