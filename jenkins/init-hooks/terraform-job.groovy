#!groovy

import jenkins.model.*
import hudson.security.*
import hudson.util.*;
import jenkins.install.*;
import jenkins.model.Jenkins
import java.io.File
import hudson.model.FreeStyleProject
import com.cloudbees.hudson.plugins.folder.*

def jobName = 'terraform'
def modmHome = System.getenv('MODM_HOME')
def jobConfigXmlPath = "${modmHome}/source/jenkins/definitions/${jobName}/config.xml"

def instance = Jenkins.get()
def xmlContent = new File(jobConfigXmlPath).text
def xmlStream = new StringBufferInputStream(xmlContent)

def job = instance.getItem(jobName)
if (job != null) {
    println "--> job ${jobName} already exists, deleting it"
    job.delete()
}

job = instance.createProjectFromXML(jobName, xmlStream)
instance.save()