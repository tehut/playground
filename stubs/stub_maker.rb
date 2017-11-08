
require 'json'

def get_file_as_string(filename)
  # filename = "/Users/tehut/test/" + filename
  data = ''
  f = File.open(filename, "r")
  f.each_line do |line|
    data += line
  end
  return data
end

class FileTree

  attr_accessor :hash

  def initialize(path)
    @hash = make_hash(path)
  end

  def to_json
    @hash.to_json
  end

  def to_html(hash=@hash, name='root')
    html = ""
    html << "<ul>"
    html << "<li>" << name << "</li>"
    if hash
      hash.each do |key, val|
        if !val
          html << "<li>" << key << "</li>"
        else
          html << "<li>" << self.to_html(val["children"], key) << "</li>"
        end
      end
    end
    html << "</ul>"
    html
  end

  def to_s
    @hash.to_s
  end

  def to_h
    @hash
  end


  private

  def make_hash(path)
    hash = {}
    Dir.entries(path).select { |file|
      !['.', '..'].include? file
    }.each do |file|
      full_path = File.join(path, file)
      if file !=  "974dd0a850880a1fd5ea088253e44f134876f2fe.yaml"
        if File.directory? full_path
          hash[file] = {
            "children" => make_hash(full_path),
            "type" => 'directory',
            "full_path" => full_path
          }
        else
          hash[file] = get_file_as_string(full_path)


        end
      else
        next
      end
    end
    hash
  end

end

dir = ARGV[0].to_s
@t = FileTree.new(dir)
print @t.to_json

