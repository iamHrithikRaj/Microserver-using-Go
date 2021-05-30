import * as express from "express"
import * as cors from "cors"

const app = express()

app.use(cors({
    origin: ["http://localhost:3000"]
}))

app.use(express.json())

const PORT = process.env.PORT || 8000

app.listen(PORT, () => console.log(`Listening on port ${PORT}`))