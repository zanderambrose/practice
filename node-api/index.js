import express from 'express'
import puppeteer from 'puppeteer'

const app = express();
const port = 3000;

app.get('/scrape', async (req, res) => {
    const browser = await puppeteer.launch({ headless: "new" });
    const page = await browser.newPage();

    await page.goto('https://tickets.smokejazz.com');

    const allEvents = await page.evaluate(() => {
        const events = Array.from(document.querySelectorAll('div.performances'))
        if (events.length) {
            const data = events.map(event => ({
                bandName: event.querySelector('h3').textContent,
                dateOfPerformance: event.querySelector('h4').textContent,
                bandInfo: event.querySelector('p').textContent
            }))

            return data
        }

        return
    })

    await browser.close();

    res.json(allEvents[0]);
});

app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});
