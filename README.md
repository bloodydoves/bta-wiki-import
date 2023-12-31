# bta-wiki-import

A cli tool to manage various aspects of BTA's [Cargo](https://www.mediawiki.org/wiki/Extension:Cargo)-enabled wiki.

## Usage

### Import data on new update

First you need to convert the defs to wikitext or "export" them:
`bta-wiki-import export ./mod-directory ./path-to-wikitext`

Then you can upload them to the wiki or "import" them:
`bta-wiki-import import -u Username@botname -l https://WEBSITE/api.php --passfile ./file-with-password ./path-to-wikitext` 

### Cache Purge to redo templating across the site

`bta-wiki-import cache-purge -u https://WEBSITE/api.php -l https://WEBSITE/files/list_of_mechs.txt`

The `list_of_mechs.txt` file is maintained out of band by wiki Admins, if you think it is out of date, ask them to update it.

### Bulk Image upload

In a directory with files named as they are expected on the wiki:

`bta-wiki-import bulk-upload -u Username@botname -l https://WEBSITE/api.php --passfile ./file-with-password`

## Contributing

1. Please limit your branches and code changes to 1 feature and its related work.
2. Cut your PR. Codeowners are automatically added to it.
3. Respond to any feedback.
4. Merge on approval.
5. Tagging  
    a. If new code was pushed, push a new tag to `master`. Tags must follow [semantic versioning](https://semver.org/) in the release format `v$major.$minor.$patch`.  
    b. If no new code was pushed, i.e. updating this README, no new tags are required.  
6. Go releaser will automatically build a new release. Check the Actions page for any errors. Fix or report them as necessary.  
