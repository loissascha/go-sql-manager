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
        <div className="w-full h-full grid-cols-[auto_1fr] bg-base text-text">
            <div>Top</div>
            <div>Hello w</div>
        </div>
    );
}

export default App;
