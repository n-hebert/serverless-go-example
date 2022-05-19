# Sample Go & Serverless

## What is this project?

I copied, updated and adapted for Canada the example files at [mthenw/serverles-go-plugin](https://github.com/mthenw/serverless-go-plugin).

✓ - A proven capacity to google & ~~steal~~ repurpose OSS code -- 100/10 technical skills

## Project Usage - Deploying the Lambda

1. Clone
2. Install [Serverless Framework](serverless.com) via npm
3. Run `npm install` in the package directory to get the plugin
4. Run `sls deploy` from within this folder

You will get an "example" Go plugin Serverless framework project.
And the junk that comes with it.

## Lambda Usage - Basics

### Option A: ClickOps Plebe Tier
You can head to the [Lambda Console page for this function](https://ca-central-1.console.aws.amazon.com/lambda/home?region=ca-central-1#/functions/example-service-dev-example) and run it from the "Test" tab.

But keep reading if you want to TURBO.

### Option B: CLI Hackgod TURBO

Open that ~hAcKer^MoDe= command line and run either
```
serverless invoke -f example
```

Smooth as butter. Or more for those who do not disregard tradition, the more ascetic fashion is as such:

```
$ aws lambda invoke --function-name example-service-dev-example response.txt
```

Now you can't see all those lovely GUIs AWS worked so hard on.
Checkmate, AWStheists.
