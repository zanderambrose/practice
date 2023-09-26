import express from 'express'
import puppeteer from 'puppeteer'

const app = express();
const port = 3000;

app.get('/scrape', async (req, res) => {
    console.log('route hit')
    const browser = await puppeteer.launch({ headless: "new" });
    const page = await browser.newPage();

    await page.goto('https://tickets.smokejazz.com');

    const detailsDiv = await page.waitForSelector('div > .details');
    console.log("detailsDiv", detailsDiv)

    const htmlContent = await page.content();

    await browser.close();

    res.json({ hello: htmlContent });
});

app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});
