part of openapi.api;

class JobSub {
  
  String templateID = null;
  
  String templateName = null;
  
  String status = null;
  
  String language = null;
  
  String script = null;
  
  int exitCode = null;
  
  String stdout = null;
  
  String stderr = null;
  
  int elapseSeconds = null;
  
  int updateAt = null;
  JobSub();

  @override
  String toString() {
    return 'JobSub[templateID=$templateID, templateName=$templateName, status=$status, language=$language, script=$script, exitCode=$exitCode, stdout=$stdout, stderr=$stderr, elapseSeconds=$elapseSeconds, updateAt=$updateAt, ]';
  }

  JobSub.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    templateID = json['templateID'];
    templateName = json['templateName'];
    status = json['status'];
    language = json['language'];
    script = json['script'];
    exitCode = json['exitCode'];
    stdout = json['stdout'];
    stderr = json['stderr'];
    elapseSeconds = json['elapseSeconds'];
    updateAt = json['updateAt'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (templateID != null)
      json['templateID'] = templateID;
    if (templateName != null)
      json['templateName'] = templateName;
    if (status != null)
      json['status'] = status;
    if (language != null)
      json['language'] = language;
    if (script != null)
      json['script'] = script;
    if (exitCode != null)
      json['exitCode'] = exitCode;
    if (stdout != null)
      json['stdout'] = stdout;
    if (stderr != null)
      json['stderr'] = stderr;
    if (elapseSeconds != null)
      json['elapseSeconds'] = elapseSeconds;
    if (updateAt != null)
      json['updateAt'] = updateAt;
    return json;
  }

  static List<JobSub> listFromJson(List<dynamic> json) {
    return json == null ? List<JobSub>() : json.map((value) => JobSub.fromJson(value)).toList();
  }

  static Map<String, JobSub> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, JobSub>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = JobSub.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of JobSub-objects as value to a dart map
  static Map<String, List<JobSub>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<JobSub>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = JobSub.listFromJson(value);
       });
     }
     return map;
  }
}

