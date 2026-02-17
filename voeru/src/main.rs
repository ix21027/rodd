use std::path::Path;

use chaser_oxide::browser::{Browser, BrowserConfig};
use chaser_oxide::fetcher::{BrowserFetcher, BrowserFetcherOptions};
use futures::StreamExt;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Fetcher browser
    let download_path = Path::new("./download");
    tokio::fs::create_dir_all(&download_path).await?;
    let fetcher = BrowserFetcher::new(
        BrowserFetcherOptions::builder()
            .with_path(&download_path)
            .build()?,
    );
    let info = fetcher.fetch().await?;
   
    println!("it worked!");
    Ok(())

}
