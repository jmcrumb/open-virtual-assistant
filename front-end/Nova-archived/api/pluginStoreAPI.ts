export class Review {
    id: string;
    sourceReview: string;
    user: string;
    plugin: string;
    rating: number;
    content: string;

    constructor(json: {[key: string]: any}) {
        this.id = json.id;
        this.sourceReview = json.sourceReview;
        this.user = json.user;
        this.plugin = json.plugin;
        this.rating = json.rating;
        this.content = json.content;
    }
}

export class Report {
    id: string;
    user: string;
    plugin: string;
    content: string;
    is_resolved: string;

    constructor(json: {[key: string]: any}) {
        this.id = json.id;
        this.user = json.user;
        this.plugin = json.plugin;
        this.content = json.content;
        this.is_resolved = json.isResolved;
    }
}

export class Plugin {
    id: string;
    publisher: string;
    sourceLink: string;
    about: string;
    downloadCount: number;
    publishedOn: Date;

    constructor(json: {[key: string]: any}) {
        this.id = json.id;
        this.publisher = json.publisher;
        this.sourceLink = json.sourceLink;
        this.about = json.about;
        this.downloadCount = json.downloadCount;
        this.publishedOn = json.publishedOn;
    }
}

export class PluginStoreAPI {

}