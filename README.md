# bta-wiki-import

A cli tool to dump various HBS BT "defs" to wikitext.

## Usage

First you need to convert the defs to wikitext or "export" them:
`bta-wiki-import export ./mod-directory ./path-to-wikitext`

Then you can upload them to the wiki or "import" them:
`bta-wiki-import import -u Username@botname -l https://WEBSITE/api.php --passfile ./file-with-password ./path-to-wikitext` 

## Contributing

1. Please limit your branches and code changes to 1 feature and its related work.
2. Cut your PR. Codeowners are automatically added to it.
3. Respond to any feedback.
4. Merge on approval.
5. Tagging  
    a. If new code was pushed, push a new tag to `master`. Tags must follow [semantic versioning](https://semver.org/) in the release format `v$major.$minor.$patch`.
    b. If no new code was pushed, i.e. updating this README, no new tags are required.
6. Go releaser will automatically build a new release. Check the Actions page for any errors. Fix or report them as necessary.
