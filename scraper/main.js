const puppeteer = require('puppeteer');
const fs = require('fs').promises;

async function scrape_json_data(browser, link) {
    const page = await browser.newPage();
    await page.goto(link);
    const js = await page.evaluate(() => {
        var { detail } = __INITIAL_STATE__;
        return detail;
    });
    return js;
}

function clean_obj(o) {
    var co = {
        found: (o.found ? o.found : null),
        listingId: (o.listingId ? o.listingId : null),
        listingAgencyReference: (o.listingAgencyReference ? o.listingAgencyReference : null),
        isSoldProperty: (o.isSoldProperty ? o.isSoldProperty : null),
        region: (o.geo.region ? o.geo.region : null),
        cityName: (o.geo.cityName ? o.geo.cityName : null),
        lon: (parseFloat(o.geo.lon) ? parseFloat(o.geo.lon) : null),
        lat: (parseFloat(o.geo.lat) ? parseFloat(o.geo.lat) : null),
        price: (o.price ? o.price : null),
        chargesPrice: (o.chargesPrice ? o.chargesPrice : null),
        caution: (o.transaction.price.caution ? o.transaction.price.caution : null),
        agency_fee: (o.transaction.price.agency_fee ? o.transaction.price.agency_fee : null),
        propertySubType: (o.propertySubType ? o.propertySubType : null),
    
        // Publisher info
        publisher_id: (o.publisher.id ? o.publisher.id : null),
        publisher_remote_visit: (o.publisher.is_allowing_remote_visit ? o.publisher.is_allowing_remote_visit : null),
        publisher_phone: (o.publisher.phone1 ? o.publisher.phone1 : null),
        publisher_name: (o.publisher.name ? o.publisher.name : null),
        publisher_athome_id: (o.publisher.athome_id ? o.publisher.athome_id : null),
    
        // House Data
        propertySurface: (o.propertySurface ? o.propertySurface : null),
        buildingYear: (o.buildingYear ? o.buildingYear : null),
        floorNumber: (o.floorNumber ? o.floorNumber : null),
        bathroomsCount: (o.bathroomsCount ? o.bathroomsCount : null),
        bedroomsCount: (o.bedroomsCount ? o.bedroomsCount : null),
        balconiesCount: (o.balconiesCount ? o.balconiesCount : null),
        garagesCount: (o.garagesCount ? o.garagesCount : null),
        carparkCount: (o.carparksCount ? o.carparksCount : null),
        hasLivingroom: (o.hasLivingroom ? o.hasLivingroom : null),
        hasKitchen: (o.has_equipped_kitchen ? o.has_equipped_kitchen : null),
        media: []
    };

    if (o.transaction.availability != null) {
        co.availability = (o.transaction.availability.text ? o.transaction.availability.text : null);
    }

    if (o.description != null) {
        co.description = (o.description.fr ? o.description.fr : null);
        if (o.description.en !== "") {
            co.description = o.description.en;
        }
    }

    for (var i in o.media.items) {
        const item = o.media.items[i];
        co.media.push('https://i1.static.athome.eu/images/annonces2/image_' + item.uri);
    }

    return co;
}

(async () => {
    const browser = await puppeteer.launch({ headless: true });
    const data = await fs.readFile(process.env.LINKS_PATH, "utf-8");
    const links = data.split(/\r?\n/);
    
    const linkPromises = links.map(async (link) => {
        let id_pattern = /id-(\d+)/;
        let id = link.match(id_pattern)[1];

        console.log("scraping: " + id);
        var obj = await scrape_json_data(browser, link);

        if (obj.found) {
            obj = clean_obj(obj);
            obj.link = link;
            return obj;
        }
        
        return { found: false, listingId: parseInt(id) };
    });
    
    const objs = await Promise.all(linkPromises);
    
    await fs.writeFile(process.env.JSON_OUT, JSON.stringify(objs));
    await browser.close()
})();
